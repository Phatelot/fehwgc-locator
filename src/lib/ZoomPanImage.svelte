<script lang="ts">
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

	let isDragging: boolean = $state(false);
	let isPinching: boolean = $state(false);
	let startX: number = $state(0);
	let startY: number = $state(0);
	let startDist: number = $state(0);
	let startScale: number = $state(1);
	let pinchCenterX: number = $state(0);
	let pinchCenterY: number = $state(0);

	function startDrag(event: PointerEvent | TouchEvent) {
		if (event instanceof TouchEvent && event.touches.length === 2) {
			isDragging = false;
			isPinching = true;
			startDist = getDistance(event.touches);
			startScale = scale;
			[pinchCenterX, pinchCenterY] = getCenter(event.touches);
			return;
		}

		isDragging = true;
		const e = (event instanceof TouchEvent ? event.touches[0] : event)!;
		startX = e.clientX - x;
		startY = e.clientY - y;
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
		x = e.clientX - startX;
		y = e.clientY - startY;
		clampPosition();
		event.preventDefault();
	}

	function endDrag() {
		isDragging = false;
		isPinching = false;
	}

	function onWheel(event: WheelEvent) {
		let newScale = scale + event.deltaY * -0.001;
		changeScale(newScale, event.clientX, event.clientY);
		event.preventDefault();
	}

	function changeScale(newScale: number, centerX: number, centerY: number) {
		const oldScale = scale;

		const rect = container.getBoundingClientRect();
		maxScale = imgEl.naturalWidth / container.getBoundingClientRect().width * 2
		newScale = clampNumber(newScale, minScale, maxScale);
		const cx = centerX - rect.left - rect.width/2;
		const cy = centerY - rect.top - rect.height/2;

		const imgCenterX = cx - x;
		const imgCenterY = cy - y;

		const scaleRatio = newScale / oldScale;
		x -= imgCenterX * (scaleRatio - 1);
		y -= imgCenterY * (scaleRatio - 1);
		scale = newScale;

		clampPosition();
	}

	function clampPosition() {
		if (!imgEl || !container) return;

		const viewWidth = container.clientWidth;
		const viewHeight = container.clientHeight;

		const minX = viewWidth*(1-scale)/2;
		const minY = viewHeight*(1-scale)/2;

		const maxX = -minX;
		const maxY = -minY;

		x = clampNumber(x, minX, maxX);
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
		onload={clampPosition}
		style="transform: translate({x}px, {y}px) scale({scale});"
	/>
</div>

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
