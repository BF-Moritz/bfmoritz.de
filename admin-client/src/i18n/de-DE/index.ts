import pages from './pages/index';

export default {
	failed: 'Aktion fehlgeschlagen',
	success: 'Aktion war erfolgreich',
	layouts: {
		main: {
			links: {
				titles: {
					community: 'Community',
					socials: 'Soziale Medien',
					code: 'Code',
				},
				community: {
					discord: 'Discord',
					ts: 'Teamspeak',
				},
				socials: {
					yt: 'YouTube',
					twitch: 'Twitch',
					instagram: 'Instagram',
					twitter: 'Twitter',
				},
				code: {
					gitlab: 'Gitlab',
					github: 'Github',
				},
				settings: {
					title: 'Einstellungen',
					open: 'Einstellungen öffnen',
				},
			},
		},
	},
	routes: {
		home: 'Home',
		about: 'Über Mich',
		gear: 'Equipment',
		videos: 'Videos',
	},
	pages,
};
