// Temporary placeholder shown wherever the old "pick any players" experience used to be.
// Removed once the pack / collection / squad phases ship.
export default function FantasyRebuildNotice({ compact = false }: { compact?: boolean }) {
  return (
    <div className={`chamfer border border-bone/10 bg-char-2 text-center ${compact ? "p-6" : "p-10"}`}>
      <p className="font-stat text-[10px] uppercase tracking-[0.3em] text-ember">Fantasy is being rebuilt</p>
      <p className="mt-3 font-display text-3xl font-black uppercase leading-none md:text-4xl">
        Packs, cards &amp; squads are coming
      </p>
      <p className="mx-auto mt-4 max-w-xl text-sm leading-6 text-ash">
        Fantasy is moving to a new format: open packs, collect player cards and build your squad from the cards you own.
        Picking players freely from the full player list is no longer available. Your team is saved and will carry over.
      </p>
    </div>
  );
}
