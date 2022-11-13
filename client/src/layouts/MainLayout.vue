<template>
	<q-layout view="lHh Lpr lFf">
		<q-header class="bg-primary">
			<q-toolbar>
				<q-btn
					flat
					dense
					round
					icon="menu"
					aria-label="Menu"
					@click="toggleLeftDrawer"
				/>

				<q-toolbar-title>
					{{ $router.currentRoute.value.meta.header }}
				</q-toolbar-title>
				<q-btn
					flat
					unelevated
					stretch
					:ripple="false"
					@click="saveLang('de-DE')"
					class="q-py-none q-px-sm"
				>
					<q-img
						src="/icons/german-96x96.png"
						:img-style="{
							opacity:
								i18n.locale.value == 'de-DE' ? '100%' : '40%',
						}"
						width="45px"
					/>
					<q-tooltip anchor="center left" self="center right"
						>Sprache: Deutsch</q-tooltip
					>
				</q-btn>
				<q-btn
					flat
					unelevated
					stretch
					:ripple="false"
					@click="saveLang('en-US')"
					class="q-py-none q-px-sm"
				>
					<q-img
						src="/icons/great-britain-96x96.png"
						:img-style="{
							opacity:
								i18n.locale.value == 'en-US' ? '100%' : '40%',
						}"
						width="45px"
					/>
					<q-tooltip anchor="center left" self="center right"
						>Language: English</q-tooltip
					>
				</q-btn>
			</q-toolbar>
			<q-tabs class="q-dark text-secondary">
				<q-route-tab
					v-for="link in routeLinks"
					:key="link.name"
					:ripple="false"
					:name="link.name"
					:icon="link.icon"
					:label="$t('routes.' + link.name)"
					:to="{ name: link.name }"
				>
					<q-badge
						v-if="link.badge !== null"
						color="purple"
						text-color="white"
						floating
						>{{ link.badge }}</q-badge
					>
				</q-route-tab>
			</q-tabs>
		</q-header>

		<q-drawer v-model="leftDrawerOpen" bordered>
			<q-list>
				<q-item-label header>{{
					$t('layouts.main.links.titles.community')
				}}</q-item-label>
				<EssentialLink
					v-for="link in communityLinks"
					:key="link.title"
					:title="$t(`layouts.main.links.community.${link.title}`)"
					:caption="link.caption"
					:icon="link.icon"
					:link="link.link"
					:external="link.external"
				/>

				<q-separator />

				<q-item-label header>{{
					$t('layouts.main.links.titles.socials')
				}}</q-item-label>
				<EssentialLink
					v-for="link in socialLinks"
					:key="link.title"
					:title="$t(`layouts.main.links.socials.${link.title}`)"
					:caption="link.caption"
					:icon="link.icon"
					:link="link.link"
					:external="link.external"
				/>

				<q-separator />

				<q-item-label header>{{
					$t('layouts.main.links.titles.code')
				}}</q-item-label>
				<EssentialLink
					v-for="link in codeLinks"
					:key="link.title"
					:title="$t(`layouts.main.links.code.${link.title}`)"
					:caption="link.caption"
					:icon="link.icon"
					:link="link.link"
					:external="link.external"
				/>
			</q-list>
		</q-drawer>

		<q-page-container
			class="full-height full-width"
			style="overflow-x: hidden; max-width=100%;"
		>
			<router-view />
		</q-page-container>
	</q-layout>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import EssentialLink from 'components/EssentialLink.vue';
import { LinkType } from '../types/link';
import { useQuasar } from 'quasar';
import { useI18n } from 'vue-i18n';

const communityLinks: LinkType[] = [
	{
		title: 'discord',
		caption: 'dc.bfmoritz.de',
		icon: 'mdi-discord',
		link: 'https://dc.bfmoritz.de',
		external: true,
	},
	{
		title: 'ts',
		caption: 'ts.bfmoritz.de',
		icon: 'mdi-microphone-outline',
		link: 'https://ts.bfmoritz.de',
		external: true,
	},
];

const socialLinks: LinkType[] = [
	{
		title: 'yt',
		caption: '@bf_moritz',
		icon: 'mdi-youtube',
		link: 'https://yt.bfmoritz.de',
		external: true,
	},
	{
		title: 'twitch',
		caption: '@bf_moritz',
		icon: 'mdi-twitch',
		link: 'https://twitch.bfmoritz.de',
		external: true,
	},
	{
		title: 'instagram',
		caption: '@bf_moritz',
		icon: 'mdi-instagram',
		link: 'https://insta.bfmoritz.de',
		external: true,
	},
	{
		title: 'twitter',
		caption: '@BF_Moritz',
		icon: 'mdi-twitter',
		link: 'https://twitter.bfmoritz.de',
		external: true,
	},
];

const codeLinks: LinkType[] = [
	{
		title: 'gitlab',
		caption: 'gitlab.bfmoritz.de',
		icon: 'mdi-gitlab',
		link: 'https://gitlab.bfmoritz.de',
		external: true,
	},
	{
		title: 'github',
		caption: 'github.bfmoritz.de',
		icon: 'mdi-github',
		link: 'https://github.bfmoritz.de',
		external: true,
	},
];

const routeLinks = [
	{
		name: 'home',
		icon: 'mdi-home-outline',
		badge: null,
	},
	{
		name: 'about',
		icon: 'mdi-account-outline',
		badge: null,
	},
	{
		name: 'gear',
		icon: 'mdi-gamepad-variant-outline',
		badge: null,
	},
	{
		name: 'videos',
		icon: 'mdi-video-outline',
		badge: null,
	},
	{
		name: 'projects',
		icon: 'mdi-code-tags',
		badge: null,
	},
];

export default defineComponent({
	name: 'MainLayout',

	components: {
		EssentialLink,
	},

	setup() {
		const $q = useQuasar();
		const leftDrawerOpen = ref(false);
		const i18n = useI18n();
		i18n.locale.value = $q.localStorage.getItem('local') || 'en-US';
		$q.localStorage.set('local', i18n.locale.value);

		return {
			communityLinks,
			socialLinks,
			codeLinks,
			leftDrawerOpen,
			routeLinks,
			i18n,
			toggleLeftDrawer() {
				leftDrawerOpen.value = !leftDrawerOpen.value;
			},
			saveLang(lang: string) {
				$q.localStorage.set('local', lang);
				i18n.locale.value = lang;
			},
		};
	},
});
</script>

<style lang="scss" src="./MainLayout.scss" scoped />
