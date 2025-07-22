<script lang="ts">
    import { getPublicImageLink } from "./asset_utils";
    import CharacterSelector from "./CharacterSelector.svelte";
    import { formatCoordinate, formatScale } from "./format_utils";
    import { positionsByCharacterSlug } from "./positions";
    import ZoomPanImage from "./ZoomPanImage.svelte";

    let scale: number = $state(1);
    let x: number = $state(0);
    let y: number = $state(0);

    let selectedCharacter: string | null = $state('edelgard_broken')

    function goToSelectedCharacter() {
        const position = positionsByCharacterSlug[selectedCharacter || ''] || {
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
</script>

<div class="container">

    <div class="options">
        <CharacterSelector bind:selectedCharacterSlug={selectedCharacter} />
        <button onclick="{() => goToSelectedCharacter()}">Go to selected character</button>

        <br/>
        <p id="coordinates-p">
            {selectedCharacter} {formatScale(scale)} {formatCoordinate(x)} {formatCoordinate(y)}
        </p>
        <button onclick="{() => copyCoordinates()}">Copy coordinates</button>

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
