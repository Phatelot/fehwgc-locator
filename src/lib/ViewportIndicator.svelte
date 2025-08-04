<script lang="ts">
	let {
		x = 50,
		y = 50,
		scale = 1,
		src = "",
		size = 200 // optional override for SVG size
	} = $props();

	// Full box size in px
	const outerSize = size;

	// Derived size of the viewport square
	let viewSize = $derived(outerSize / scale);

	// Top-left corner of viewport rect
	let viewX = $derived(clamp((x / 100) * outerSize - viewSize / 2, 0, outerSize - viewSize));
	let viewY = $derived(clamp((y / 100) * outerSize - viewSize / 2, 0, outerSize - viewSize));

	function clamp(n: number, min: number, max: number) {
		return Math.max(min, Math.min(n, max));
	}
</script>

<svg
	width={outerSize}
	height={outerSize}
	viewBox={`0 0 ${outerSize} ${outerSize}`}
	style="border: 1px solid #999; background: #f8f8f8;"
>
	<!-- Full image boundary -->
	<image
		href={src}
		x="0"
		y="0"
		width={outerSize}
		height={outerSize}
		preserveAspectRatio="xMidYMid slice"
	/>

	<!-- Viewport rectangle -->
	<rect
		x={viewX}
		y={viewY}
		width={viewSize}
		height={viewSize}
		fill="rgba(0, 0, 255, 0.2)"
		stroke="blue"
		stroke-width="2"
	/>
</svg>
