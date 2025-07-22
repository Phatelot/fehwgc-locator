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

<div class="container">
    <div class="options">
        <p>scale: {formatScale(scale)}</p>
        <p>x: {formatCoordinate(x)}</p>
        <p>y: {formatCoordinate(y)}</p>
        <button onclick="{() => scaleTo2()}">scale = 2</button>
    </div>

    <div class="zoompanimage">
        <ZoomPanImage bind:scale={scale} src={getPublicImageLink("spread.webp")} bind:x={x} bind:y={y} />
    </div>
</div>

<style>
    p, button {
        color: black;
        -webkit-text-stroke-width: 0px;
    }

    .container {
        display: flex;
    }

    @media (max-width: 640px) {
        .container {
            flex-direction: column;
        }
    }

    .options {
        padding: 30px 30px 30px 0;
    }

    .zoompanimage {
        flex-grow: 1;
    }
</style>
