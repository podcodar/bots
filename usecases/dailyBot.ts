import { client } from '@repository/supabase.ts';
import {
	CreateDailyRecord,
	CreateDailyScoreboard,
	DailyRecord,
	DailyScoreboard,
} from '@entities';

export function addDailyRecord(dailyRecord: CreateDailyRecord) {
	return client.from<DailyRecord>('daily_record').insert([dailyRecord]);
}

export async function computeDailyScoreboardPoint(
	userId: string,
	name: string,
) {
	const { error, data } = await getScoreboardByUserId(userId);
	if (error != null) {
		throw error;
	}

	if (data.length === 0) {
		// if no user on the scoreboard, create one
		const dailyScoreboard: CreateDailyScoreboard = {
			userId,
			name,
			points: 1,
			currentStreak: 1,
		};

		return await upsertDailyScoreboard(dailyScoreboard);
	}

	const [dailyScoreboard] = data;

	// add 1 point to the user
	dailyScoreboard.points += 1;

	// add one extra point to the user if he had 5 points in a row
	if (await countDailyActivity(7) === 5) dailyScoreboard.points += 1;

	// increment the current streak unless no record was found on the last day
	dailyScoreboard.currentStreak += 1;
	if (await countDailyActivity(1) === 0) dailyScoreboard.currentStreak = 0;

	// save changes
	return await upsertDailyScoreboard(dailyScoreboard);
}

async function countDailyActivity(rangeDays: number) {
	// TODO: count dailies records from last week
	const today = new Date();
	const rangeDate = new Date(Date.now() - rangeDays * 24 * 60 * 60 * 1000);

	const { error, data } = await client
		.from('daily_record')
		.select('*', { count: 'exact' })
		.lt('created_at', today)
		.gte('created_at', rangeDate);

	if (error != null) {
		console.error(error);
		return 0;
	}

	return data.length;
}

function upsertDailyScoreboard(
	dailyScoreboard: CreateDailyScoreboard | DailyScoreboard,
) {
	return client.from<DailyScoreboard>('daily_scoreboard').upsert([
		dailyScoreboard,
	]);
}

function getScoreboardByUserId(userId: string) {
	return client.from<DailyScoreboard>('daily_scoreboard').select('*').eq(
		'userId',
		userId,
	);
}
