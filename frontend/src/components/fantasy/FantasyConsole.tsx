"use client";

import { useCallback, useEffect, useState } from "react";
import Link from "next/link";
import ErrorState from "@/components/ui/ErrorState";
import Skeleton from "@/components/ui/Skeleton";
import { ApiError, getMyFantasyTeam } from "@/lib/api";
import { clearSession, getToken } from "@/lib/auth";
import type { FantasyTeam } from "@/types";
import CreateFantasyTeamForm from "./CreateFantasyTeamForm";
import FantasyRebuildNotice from "./FantasyRebuildNotice";

// Entry point of the Fantasy tab: log in -> create your fantasy team -> (placeholder).
// The old "pick any 4 players from the pool" builder has been removed; the pack, collection
// and squad screens will replace the placeholder in later phases.
export default function FantasyConsole({ tournamentId }: { tournamentId: number }) {
  const [authChecked, setAuthChecked] = useState(false);
  const [loggedIn, setLoggedIn] = useState(false);
  const [team, setTeam] = useState<FantasyTeam | null | undefined>(undefined);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    setLoggedIn(!!getToken());
    setAuthChecked(true);
  }, []);

  const loadTeam = useCallback(() => {
    setError(null);
    getMyFantasyTeam(tournamentId)
      .then(setTeam)
      .catch((e) => {
        if (e instanceof ApiError && e.status === 404) {
          setTeam(null);
        } else if (e instanceof ApiError && e.status === 401) {
          clearSession();
          setLoggedIn(false);
        } else {
          setError(e instanceof Error ? e.message : "Something went wrong");
        }
      });
  }, [tournamentId]);

  useEffect(() => {
    if (loggedIn) loadTeam();
  }, [loggedIn, loadTeam]);

  if (!authChecked) {
    return (
      <div className="mx-auto max-w-7xl px-5 py-10">
        <Skeleton className="h-48" />
      </div>
    );
  }

  if (!loggedIn) {
    return (
      <div className="relative min-h-[70vh] overflow-hidden px-5 py-20">
        <div className="pointer-events-none absolute inset-0">
          <div className="absolute left-1/2 top-20 h-72 w-72 -translate-x-1/2 rounded-full bg-ember/10 blur-3xl" />
        </div>

        <div className="relative mx-auto max-w-md">
          <div className="chamfer border border-bone/10 bg-char-2/90 p-10 text-center shadow-2xl">
            <div className="mx-auto mb-6 flex h-16 w-16 items-center justify-center chamfer-sm border border-ember/30 bg-ember/10">
              <span className="font-display text-2xl font-black text-ember">FF</span>
            </div>

            <p className="font-stat text-[10px] uppercase tracking-[0.3em] text-ash">FFWS 2026 Fantasy</p>

            <p className="mt-2 font-display text-4xl font-black uppercase leading-none">Log in to play</p>

            <p className="mt-4 text-sm leading-6 text-ash">Create your fantasy team and get ready for the new card-based game.</p>

            <Link
              href="/login"
              className="mt-8 inline-flex items-center justify-center chamfer-sm border border-ember/40 bg-ember px-7 py-3 font-stat text-xs font-bold uppercase tracking-[0.18em] text-char transition hover:bg-bone"
            >
              Go to login
              <span className="ml-3 text-base">→</span>
            </Link>
          </div>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="mx-auto max-w-3xl px-5 py-16">
        <ErrorState message={error} />
      </div>
    );
  }

  if (team === undefined) {
    return (
      <div className="mx-auto max-w-7xl px-5 py-10">
        <Skeleton className="h-48" />
      </div>
    );
  }

  if (team === null) {
    return (
      <div className="relative min-h-[70vh] overflow-hidden px-5 py-12">
        <div className="pointer-events-none absolute inset-0">
          <div className="absolute right-0 top-20 h-80 w-80 rounded-full bg-ember/10 blur-3xl" />
          <div className="absolute bottom-0 left-0 h-72 w-72 rounded-full bg-bone/5 blur-3xl" />
        </div>

        <div className="relative mx-auto max-w-5xl">
          <div className="mb-10">
            <p className="font-stat text-[10px] uppercase tracking-[0.3em] text-ember">FFWS 2026 Fantasy</p>

            <h1 className="mt-2 font-display text-5xl font-black uppercase leading-none tracking-tight md:text-6xl">
              Create your
              <span className="block text-ember">fantasy team</span>
            </h1>

            <p className="mt-5 max-w-2xl text-sm leading-6 text-ash md:text-base">
              Choose your identity. Your team will hold your card collection and squads once packs open.
            </p>
          </div>

          <CreateFantasyTeamForm tournamentId={tournamentId} onCreated={loadTeam} />
        </div>
      </div>
    );
  }

  return (
    <div className="mx-auto max-w-3xl space-y-6 px-5 py-10">
      <div className="chamfer flex items-center gap-4 border border-bone/10 bg-char-2 p-6">
        <div className="flex h-14 w-14 items-center justify-center chamfer-sm border border-ember/25 bg-ember/10">
          <span className="font-display text-xl font-black text-ember">{team.team_name.slice(0, 2).toUpperCase()}</span>
        </div>
        <div>
          <p className="font-stat text-[10px] uppercase tracking-widest text-ash">Your fantasy team</p>
          <p className="font-display text-3xl font-black uppercase leading-none">{team.team_name}</p>
        </div>
      </div>

      <FantasyRebuildNotice />
    </div>
  );
}
