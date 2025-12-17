<script lang="ts">
    import { onMount, tick } from "svelte";
    import { getPublicImageLink } from "./asset_utils";
    import CharacterSelector from "./CharacterSelector.svelte";
    import { formatCoordinate, formatScale } from "./format_utils";
    import { getRandomLocatedOutfitSlug, positionsByCharacterSlug } from "./positions";
    import ZoomPanImage from "./ZoomPanImage.svelte";
    import ViewportIndicator from "./ViewportIndicator.svelte";
    import { getSlugsFromCombinedSlug } from "./metadata_utils";
    import Chibi from "./Chibi.svelte";

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
        if (!positionsByCharacterSlug[selectedCharacterSlug || '']) {
            return
        }
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

    function goToOutfitPopup() {
        if (!selectedCharacterSlug) {
            return;
        }
        const slug = getSlugsFromCombinedSlug(selectedCharacterSlug)

        const state = retrieveStateFromLocalStorage();
        state.page = "OUTFIT_CHART";
        state.selectedCharacterSlug = slug?.characterSlug;
        state.selectedOutfitSlug = slug?.outfitSlug;
        localStorage.setItem("fehwgc", JSON.stringify(state));
        window.location.href = 'https://phatelot.github.io/fehwgc';
    }

    function retrieveStateFromLocalStorage(): any {
      const retrieved = localStorage.getItem("fehwgc");
      if (!retrieved) {
        return {}
      }
      return JSON.parse(retrieved);
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

        <button onclick="{() => goToOutfitPopup()}">Read her detailed info</button>
        <button onclick="{() => goToRandomCharacter()}">Go to random character</button>

        <div class="smallimagescontainer">
            <Chibi {selectedCharacterSlug}/>
            <ViewportIndicator {x} {y} {scale} src={getPublicImageLink("spread.webp")}/>
        </div>

        {#if true}
            <p id="coordinates-p">
                "{selectedCharacterSlug}": &lbrace;scale: {formatScale(scale)}, x: {formatCoordinate(x)}, y: {formatCoordinate(y)}&rbrace;,
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

    button {
        margin-bottom: 10px;
        width: 100%;
    }

    .container {
        display: flex;
    }

    .options {
        padding: 30px 30px 30px 0;
        max-width: 500px;
    }

    @media (max-width: 1080px) {
        .container {
            flex-direction: column;
            align-items: center;
        }

        .options {
            padding-right: 0px;
        }

        .zoompanimage {
            width: 100%;
        }
    }

    .zoompanimage {
        flex-grow: 1;
    }

    .smallimagescontainer {
        display: flex;
        justify-content: space-between;
    }
</style>
