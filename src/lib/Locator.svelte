<script lang="ts">
    import { getPublicImageLink } from "./asset_utils";
    import { formatCoordinate, formatScale } from "./format_utils";
    import ZoomPanImage from "./ZoomPanImage.svelte";

    let scale: number = $state(1);
    let x: number = $state(0);
    let y: number = $state(0);

    function scaleTo2() {
        scale = 2;
    }

</script>

<ZoomPanImage bind:scale={scale} src={getPublicImageLink("spread.webp")} bind:x={x} bind:y={y} />
<p>scale: {formatScale(scale)}</p>
<p>x: {formatCoordinate(x)}</p>
<p>y: {formatCoordinate(y)}</p>
<button onclick="{() => scaleTo2()}">scale = 2</button>

<style>
    p {
        color: black;
    }
</style>
