import { baseMetadata } from "./metadata"

export function getSlugsFromCombinedSlug(selectedCharacterSlug: string): {
	characterSlug: string,
	outfitSlug: string,
} | undefined {
	return baseMetadata
		.flatMap(game => game.characters)
		.filter(character => selectedCharacterSlug.startsWith(character.nameSlug))
		.flatMap(character =>
			[...character.outfits.map(outfit => outfit.outfitSlug), 'broken'].map(outfitSlug => ({
				characterSlug: character.nameSlug,
				outfitSlug,
			})),
		)
		.filter(e => selectedCharacterSlug === `${e.characterSlug}_${e.outfitSlug}`)[0]
}
