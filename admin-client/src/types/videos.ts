export type VideoType = {
	id: string;
	url: string;
	title: string;
	description: string;
	channel_id: string;
	channel_title: string;
	playlist_id: string;
	published_at: string;
	thumbnail: {
		url: string;
		width: number;
		height: number;
	};
	tags: string[];
	meta: null | MetaType;
};

export type MetaType = {
	description: LanguageStringType | null;
	correction: LanguageStringType | null;
	text_sources: LanguageStringType[];
	video_sources: LanguageStringType[];
	audio_sources: LanguageStringType[];
};

export type LanguageStringType = {
	'de-DE': string;
	'en-US': string;
};
