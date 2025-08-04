<script lang="ts">
    import { baseMetadata } from "./metadata";
    import { positionsByCharacterSlug } from "./positions";

	let {
		selectedCharacterSlug = $bindable('edelgard_broken')
	}:
	{
		selectedCharacterSlug: string | null
	} = $props();

	const choices : {
		text: string,
		value: string,
	}[] = baseMetadata
		.flatMap(game => game.characters)
		.flatMap(character => {
			const outfits = character.outfits.map(outfit => ({
				text: `${character.name} - ${outfit.outfit.toLowerCase()}`,
				value: `${character.nameSlug}_${outfit.outfitSlug}`
			}))
			outfits.push({
				text: `${character.name} - broken`,
				value: `${character.nameSlug}_broken`
			})
			return outfits;
		})
		.filter(outfit => Object.hasOwn(positionsByCharacterSlug, outfit.value))
		.sort((a, b) => a.text > b.text ? 1 : -1)


	function updateCsParam(newSlug: string) {
		const hash = location.hash.slice(1) || '/';
		const [path, queryString] = hash.split('?');
		const searchParams = new URLSearchParams(queryString || '');

		searchParams.set('cs', newSlug);

		// Replace the hash silently without triggering 'hashchange'
		// by using history.replaceState
		const newHash = path + '?' + searchParams.toString();
		history.replaceState(null, '', '#' + newHash);
	}

</script>

<form>
    <select bind:value={selectedCharacterSlug} style="font-weight: 600;" onchange={(e) => updateCsParam(e.currentTarget.value)}>
      {#each choices as choice}
        <option style="font-weight: 600;" value={choice.value}>
         {choice.text}
        </option>
      {/each}
    </select>
</form>

<style>
  form {
    display: flex;
  }

  @media (width <= 768px) {
    form {
      flex-direction: column;
    }
  }

  select, option {
    -webkit-text-stroke-width: 0px;
  }

</style>
