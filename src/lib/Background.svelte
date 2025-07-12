<script lang="ts">
    import { getPublicImageLink } from "./asset_utils";
</script>

<svg viewBox="0 0 5432 2387" xmlns="http://www.w3.org/2000/svg" class="bg">
	<defs>
		<linearGradient id="bg-mask-gradient" x1="0" x2="0" y1="0" y2="1">
		  <stop offset="0%" stop-color="#ffffff"></stop>
		  <stop offset="60%" stop-color="#ffffff"></stop>
		  <stop offset="100%" stop-color="#000000"></stop>
		</linearGradient>
		<mask id="bg-mask" x="0%" y="0%" width="100%" height="100%" maskUnits="userSpaceOnUse">
		  <rect x="0%" y="0%" width="100%" height="100%" fill="url(#bg-mask-gradient)" />
		</mask>
	  </defs>

	  <image xlink:href="{getPublicImageLink("BG.webp")}" width=100% mask="url(#bg-mask)"/>
</svg>

<style>

.bg {
	position: fixed;
	inset: 0;
	width: 100%;
	z-index: -1;
	pointer-events: none;
}

</style>
