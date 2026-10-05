"use client";

import Link from "next/link";
import { useEffect, useState } from "react";
import { getMyFantasyTeam } from "@/lib/api";
import { getToken } from "@/lib/auth";
import FantasyRebuildNotice from "./FantasyRebuildNotice";

// Public page of a fantasy team. It only shows the team's identity for now: squads and card
// collections come with the later phases of the card-based fantasy game.
export default function FantasyTeamViewer({
  tournamentId,
  fantasyTeamId,
  teamName,
}: {
  tournamentId: number;
  fantasyTeamId: number;
  teamName: string;
}) {
  const [isMine, setIsMine] = useState(false);

  useEffect(() => {
    if (!getToken()) return;
    let cancelled = false;
    getMyFantasyTeam(tournamentId)
      .then((mine) => {
        if (!cancelled) setIsMine(mine.id === fantasyTeamId);
      })
      .catch(() => {
        if (!cancelled) setIsMine(false);
      });
    return () => {
      cancelled = true;
    };
  }, [tournamentId, fantasyTeamId]);

  return (
    <div className="mx-auto max-w-3xl space-y-6 px-5 py-10">
      <div className="chamfer flex flex-wrap items-center justify-between gap-4 border border-bone/10 bg-char-2 p-6">
        <div>
          <p className="font-stat text-[10px] uppercase tracking-widest text-ash">{isMine ? "Your fantasy team" : "Viewing"}</p>
          <p className="font-display text-3xl font-black uppercase">{teamName}</p>
        </div>
        <Link href="/fantasy/leaderboard" className="font-stat text-xs uppercase tracking-widest text-ember hover:underline">
          Leaderboard →
        </Link>
      </div>

      <FantasyRebuildNotice compact />
    </div>
  );
}
