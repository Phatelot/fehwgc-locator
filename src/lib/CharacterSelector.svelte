<script lang="ts">
    import { baseMetadata } from "./metadata";

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

</script>

<form>
    <select bind:value={selectedCharacterSlug} style="font-weight: 600;">
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
