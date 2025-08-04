<script lang="ts">
    import { onMount, tick } from "svelte";
    import { getPublicImageLink } from "./asset_utils";
    import CharacterSelector from "./CharacterSelector.svelte";
    import { formatCoordinate, formatScale } from "./format_utils";
    import { getRandomLocatedOutfitSlug, positionsByCharacterSlug } from "./positions";
    import ZoomPanImage from "./ZoomPanImage.svelte";

    let {
		selectedCharacterSlug = $bindable('edelgard_broken')
	}:
	{
		selectedCharacterSlug?: string | null
	} = $props();

    let scale: number = $state(1);
    let x: number = $state(0);
    let y: number = $state(0);

    function goToRandomCharacter() {
        selectedCharacterSlug = getRandomLocatedOutfitSlug()
        goToSelectedCharacter()
    }

    function goToSelectedCharacter() {
        const position = positionsByCharacterSlug[selectedCharacterSlug || ''] || {
            scale: 1,
            x: 0,
            y: 0,
        }

        scale = position.scale * 1.0001
        scale = position.scale
        x = position.x
        y = position.y
    }

    function copyCoordinates() {
		const p = document.getElementById("coordinates-p");
		if (p && p.textContent) {
			navigator.clipboard.writeText(p.textContent)
				.then(() => {
					console.log("Copied to clipboard:", p.textContent);
				})
				.catch(err => {
					console.error("Failed to copy text:", err);
				});
		}
	}

    onMount(async () => {
        if (!selectedCharacterSlug) {
            return;
        }
        if (!positionsByCharacterSlug[selectedCharacterSlug]) {
            selectedCharacterSlug = '';
            return;
        }
        await tick();
        goToSelectedCharacter();
    });


</script>

<div class="container">

    <div class="options">
        <CharacterSelector bind:selectedCharacterSlug={
            () => selectedCharacterSlug,
            (scs) => {
                selectedCharacterSlug = scs;
                goToSelectedCharacter();
            }}
        />

        <br/>

        <button onclick="{() => goToRandomCharacter()}">Go to random character</button>

        {#if false}
            <p id="coordinates-p">
                {selectedCharacterSlug} {formatScale(scale)} {formatCoordinate(x)} {formatCoordinate(y)}
            </p>
            <button onclick="{() => copyCoordinates()}">Copy coordinates</button>
        {/if}
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
