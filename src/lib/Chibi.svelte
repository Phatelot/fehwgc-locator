<script lang="ts">
    import { getBrokenOutfit, getSlugsFromCombinedSlug } from "./metadata_utils";


	let {
		selectedCharacterSlug = null
	}: {
		selectedCharacterSlug: string | null
	} = $props();

	let imgSrc = $derived.by(() => {
		if (!selectedCharacterSlug) {
			return undefined;
		}
		const {characterSlug, outfitSlug} = getSlugsFromCombinedSlug(selectedCharacterSlug) || {};
		if (!(characterSlug && outfitSlug)) {
			return undefined;
		}

		let chibi = outfitSlug + "_chibi";
		if (outfitSlug == 'broken') {
			chibi = (getBrokenOutfit(characterSlug) || '') + "_chibi";
			if (['edelgard', 'ena', 'heiorun', 'hraesvelgr', 'niohoggr'].includes(characterSlug)) {
				chibi += "_alt"
			}
		}
		chibi = chibi.replace("base_", "");
		return "https://phatelot.github.io/fehwgc/characters/" + characterSlug + "_" + chibi + ".webp"
	});

</script>

{#if imgSrc}
	<img src={imgSrc} alt="chibi art of the selected character"/>
{/if}
