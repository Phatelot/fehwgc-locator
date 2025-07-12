<script lang="ts">
	let container: HTMLDivElement;
	let imgEl: HTMLImageElement;

	let {
	  src = '',
	  x = 0,
	  y = 0,
	  scale = 1
	}: {
	  src: string;
	  x: number;
	  y: number;
	  scale: number;
	} = $props();

	let isDragging: boolean = $state(false);
	let startX: number = $state(0);
	let startY: number = $state(0);

	function startDrag(event: PointerEvent | TouchEvent) {
	  isDragging = true;
	  const e = (event instanceof TouchEvent ? event.touches[0] : event)!;
	  startX = e.clientX - x;
	  startY = e.clientY - y;
	  event.preventDefault();
	}

	function onDrag(event: PointerEvent | TouchEvent) {
	  if (!isDragging) return;
	  const e = (event instanceof TouchEvent ? event.touches[0] : event)!;
	  x = e.clientX - startX;
	  y = e.clientY - startY;
	  clampPosition();
	  event.preventDefault();
	}

	function endDrag() {
	  isDragging = false;
	}

	function onWheel(event: WheelEvent) {
	  const oldScale = scale;
	  scale += event.deltaY * -0.001;
	  scale = Math.min(Math.max(0.5, scale), 4);

	  const rect = container.getBoundingClientRect();
	  const offsetX = event.clientX - rect.left - container.clientWidth / 2;
	  const offsetY = event.clientY - rect.top - container.clientHeight / 2;

	  x -= offsetX * (scale - oldScale) / oldScale;
	  y -= offsetY * (scale - oldScale) / oldScale;

	  clampPosition();
	  event.preventDefault();
	}

	function clampPosition() {
	  if (!imgEl || !container) return;

	  const imgWidth = imgEl.naturalWidth * scale;
	  const imgHeight = imgEl.naturalHeight * scale;

	  const viewWidth = container.clientWidth;
	  const viewHeight = container.clientHeight;

	  const maxX = Math.max(0, (imgWidth - viewWidth) / 2);
	  const maxY = Math.max(0, (imgHeight - viewHeight) / 2);

	  x = Math.min(Math.max(x, -maxX), maxX);
	  y = Math.min(Math.max(y, -maxY), maxY);
	}
  </script>

  <div
	bind:this={container}
	class="viewport"
	onpointerdown={startDrag}
	onpointermove={onDrag}
	onpointerup={endDrag}
	onpointercancel={endDrag}
	onpointerleave={endDrag}
	onwheel={onWheel}
  >
	<img
	  bind:this={imgEl}
	  src={src}
	  alt="Zoomable view"
	  draggable="false"
	  onload={clampPosition}
	  style="transform: translate({x}px, {y}px) scale({scale});"
	/>
  </div>

  <div class="controls">
	Zoom: <input
	  type="range"
	  min="0.5"
	  max="4"
	  step="0.01"
	  bind:value={scale}
	/>
  </div>

  <style>
	.viewport {
	  width: 100%;
	  max-width: 600px;
	  height: 400px;
	  overflow: hidden;
	  position: relative;
	  border: 2px solid #ccc;
	  touch-action: none;
	  margin-bottom: 0.5rem;
	}

	img {
	  user-select: none;
	  pointer-events: none;
	  position: absolute;
	  top: 0;
	  left: 0;
	  will-change: transform;
	}

	.controls {
	  text-align: center;
	}
  </style>
