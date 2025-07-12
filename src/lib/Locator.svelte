<script lang="ts">
    import { getPublicImageLink } from "./asset_utils";
    import ZoomPanImage from "./ZoomPanImage.svelte";
</script>

<ZoomPanImage scale={1} src={getPublicImageLink("spread.webp")} x={0} y={0} />
