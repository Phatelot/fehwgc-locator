<script lang="ts">
    import { formatCoordinate, formatScale } from "./format_utils";

	let container: HTMLDivElement;
	let imgEl: HTMLImageElement;

	const minScale = 1;
	let maxScale = $state(2);

	const borderPx = 2;

	let {
		src = "",
		x = $bindable(0),
		y = $bindable(0),
		scale = $bindable(1),
	}: {
		src: string;
		x: number;
		y: number;
		scale: number;
	} = $props();

	let innerScale = $derived(convertToInnerScale(scale)); // inverse of the linear proportion of the picture currently shown (eg if = 2, then we're showing a fourth of the image area)
	function updateOuterScale(newInnerScale: number) {
		scale = convertToOuterScale(newInnerScale)
	}
	let innerX = $derived(convertToInnerX(x)); // how many pixels the image is translated horizontally, *after* scaling (when = 0, the viewport is horizontally centered on the image)
	function updateOuterX(newInnerX: number) {
		x = convertToOuterX(newInnerX)
	}
	let innerY = $derived(convertToInnerY(y)); // how many pixels the image is translated vertically, *after* scaling (when = 0, the viewport is vertically centered on the image)
	function updateOuterY(newInnerY: number) {
		y = convertToOuterY(newInnerY)
	}

	let isDragging: boolean = $state(false);
	let isPinching: boolean = $state(false);
	let startX: number = $state(0);
	let startY: number = $state(0);
	let startDist: number = $state(0);
	let startScale: number = $state(1);
	let pinchCenterX: number = $state(0);
	let pinchCenterY: number = $state(0);

	function convertToOuterScale(innerScale: number): number {
		if (!imgEl || !container) return innerScale;
		return innerScale / (imgEl.naturalWidth / container.getBoundingClientRect().width)
	}

	function convertToInnerScale(outerScale: number): number {
		if (!imgEl || !container) return outerScale;
		return outerScale * (imgEl.naturalWidth / container.getBoundingClientRect().width)
	}

	function convertToOuterX(innerX: number): number {
		if (!imgEl || !container) return 0;
		if (innerScale === 1) return 50;
		return innerX / (container.clientWidth) * (1/innerScale - 1) / (innerScale - 1) * 100 + 50;
	}

	function convertToInnerX(outerX: number): number {
		if (!imgEl || !container) return 0;
		if (innerScale === 1) return 0;
		return (outerX - 50) * (container.clientWidth) / (1/innerScale - 1) * (innerScale - 1) / 100
	}

	function convertToOuterY(innerY: number): number {
		if (!imgEl || !container) return 0;
		if (innerScale === 1) return 50;
		return innerY / (container.clientHeight) * (1/innerScale - 1) / (innerScale - 1) * 100 + 50;
	}

	function convertToInnerY(outerY: number): number {
		if (!imgEl || !container) return 0;
		if (innerScale === 1) return 0;
		return (outerY - 50) * (container.clientHeight) / (1/innerScale - 1) * (innerScale - 1) / 100
	}

	function reset() {
		x = 50
		y = 50
		scale = 1
	}

	function startDrag(event: PointerEvent | TouchEvent) {
		if (event instanceof TouchEvent && event.touches.length === 2) {
			isDragging = false;
			isPinching = true;
			startDist = getDistance(event.touches);
			startScale = innerScale;
			[pinchCenterX, pinchCenterY] = getCenter(event.touches);
			return;
		}

		isDragging = true;
		const e = (event instanceof TouchEvent ? event.touches[0] : event)!;
		startX = e.clientX - innerX;
		startY = e.clientY - innerY;
		event.preventDefault();
	}

	function onDrag(event: PointerEvent | TouchEvent) {
		if (event instanceof TouchEvent && event.touches.length === 2) {
			const newDist = getDistance(event.touches);
			changeScale(startScale * (newDist / startDist), pinchCenterX, pinchCenterY);
			clampPosition();
			event.preventDefault();
			return;
		}

		if (!isDragging) return;
		const e = (event instanceof TouchEvent ? event.touches[0] : event)!;
		updateOuterX(e.clientX - startX);
		updateOuterY(e.clientY - startY);
		clampPosition();
		event.preventDefault();
	}

	function endDrag() {
		isDragging = false;
		isPinching = false;
	}

	function onWheel(event: WheelEvent) {
		const now = Date.now();
		let newScale = innerScale + event.deltaY * -0.001;

		changeScale(newScale, event.clientX, event.clientY);
		event.preventDefault();
	}

	function changeScale(newScale: number, centerX: number, centerY: number) {
		const oldScale = innerScale;

		const rect = container.getBoundingClientRect();
		maxScale = imgEl.naturalWidth / rect.width * 2
		newScale = clampNumber(newScale, minScale, maxScale);
		const scaleRatio = newScale / oldScale;

		const offsetX = centerX - rect.left - rect.width/2;
		const offsetY = centerY - rect.top - rect.height/2;

		// Convert current outer x/y to inner (px)
		const oldInnerX = convertToInnerX(x);
		const oldInnerY = convertToInnerY(y);

		// Adjust by scale
		const newInnerX = oldInnerX - offsetX * (scaleRatio - 1);
		const newInnerY = oldInnerY - offsetY * (scaleRatio - 1);

		updateOuterX(newInnerX);
		updateOuterY(newInnerY);
		updateOuterScale(newScale);

		clampPosition();
	}

	function clampPosition() {
		if (!imgEl || !container) return;
		const minX = 1 / 2 / innerScale * 100;
		const maxX = 100 - minX;
		x = clampNumber(x, minX, maxX);

		const minY = 1 / 2 / innerScale * 100;
		const maxY = 100 - minY;
		y = clampNumber(y, minY, maxY);
	}

	function getDistance(touches: TouchList): number {
		const dx = touches[0].clientX - touches[1].clientX;
		const dy = touches[0].clientY - touches[1].clientY;
		return Math.hypot(dx, dy);
	}

	function getCenter(touches: TouchList): [number, number] {
		const x = (touches[0].clientX + touches[1].clientX) / 2;
		const y = (touches[0].clientY + touches[1].clientY) / 2;
		return [x, y];
	}

	function clampNumber(n: number, min: number, max: number) {
		return Math.min(Math.max(n, min), max);
	}
</script>

<div
	bind:this={container}
	class="viewport"
	style="border: {borderPx}px solid #ccc;"
	onpointerdown={startDrag}
	onpointermove={onDrag}
	onpointerup={endDrag}
	onpointercancel={endDrag}
	onpointerleave={endDrag}
	onwheel={onWheel}
	ontouchstart={startDrag}
	ontouchmove={onDrag}
	ontouchend={endDrag}
	ontouchcancel={endDrag}
>
	<img
		bind:this={imgEl}
		{src}
		alt="Zoomable view"
		draggable="false"
		onload={() => {clampPosition(); reset();}}
		style="transform: translate({innerX}px, {innerY}px) scale({innerScale});"
	/>
</div>

<p>innerScale: {formatScale(1/innerScale)}</p>
<p>x: {formatCoordinate(x)}</p>
<p>y: {formatCoordinate(y)}</p>

<style>
	.viewport {
		width: 100%;
		max-width: 800px;
		aspect-ratio: 1 / 1;
		overflow: hidden;
		position: relative;
		touch-action: none;
		margin-bottom: 0.5rem;
		margin: auto;
	}

	img {
		user-select: none;
		pointer-events: none;
		position: absolute;
		top: 0;
		left: 0;
		will-change: transform;
	}
</style>
