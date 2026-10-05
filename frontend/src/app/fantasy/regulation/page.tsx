import RegulationView, { type RuleSection } from "@/components/fantasy/RegulationView";

export const metadata = { title: "Regulation · Fantasy · FFWS 2026" };

// The old rules (pick 4 players from the full pool, $100 budget, chips) no longer apply.
// The new rules will be written once the pack, collection and squad phases are live.
const sections: RuleSection[] = [
  {
    title: "Fantasy is being rebuilt",
    content: (
      <>
        <p>
          Fantasy is moving to a new format: open packs, collect player cards and build your squad from the cards you own.
          Choosing players freely from the full player list is no longer possible.
        </p>
        <p>The full regulation will be published here when the new game opens.</p>
      </>
    ),
  },
  {
    title: "What stays the same",
    content: (
      <ul className="list-disc space-y-1 pl-5">
        <li>You need a fantasy team (team name and country). It is shared with Prediction.</li>
        <li>Your fantasy team is kept and will carry over to the new game.</li>
        <li>Prediction is unchanged: see its own Regulation tab.</li>
      </ul>
    ),
  },
];

export default function FantasyRegulationPage() {
  return (
    <RegulationView
      eyebrow="Fantasy"
      title="Regulation"
      intro="The rules of the new card-based fantasy game will be published here."
      sections={sections}
    />
  );
}
