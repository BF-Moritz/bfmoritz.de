import { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
	{
		path: '/',
		component: () => import('layouts/MainLayout.vue'),
		children: [
			{
				name: 'home',
				path: '',
				component: () => import('pages/Index/IndexPage.vue'),
			},
		],
		meta: { title: 'Home', header: 'Home' },
	},
	{
		path: '/about',
		component: () => import('layouts/MainLayout.vue'),
		children: [
			{
				name: 'about',
				path: '',
				component: () => import('pages/About/AboutPage.vue'),
			},
		],
		meta: { title: 'About', header: 'About' },
	},
	{
		path: '/gear',
		component: () => import('layouts/MainLayout.vue'),
		children: [
			{
				name: 'gear',
				path: '',
				component: () => import('pages/Gear/GearPage.vue'),
			},
		],
		meta: { title: 'Gear', header: 'Gear' },
	},
	{
		path: '/videos',
		component: () => import('layouts/MainLayout.vue'),
		children: [
			{
				name: 'videos',
				path: '',
				component: () => import('pages/Videos/VideosPage.vue'),
			},
		],
		meta: { title: 'Videos', header: 'Videos' },
	},

	// Always leave this as last one,
	// but you can also remove it
	{
		path: '/:catchAll(.*)*',
		component: () => import('pages/Error/ErrorNotFound.vue'),
		meta: { title: '404', header: '404' },
	},
];

export default routes;
