<template>
	<q-card style="width: 80em; max-width: 100vw; max-height: 80vh">
		<q-card-section style="height: 45em">
			<iframe
				width="100%"
				height="100%"
				:src="`https://www.youtube-nocookie.com/embed/${video.id}`"
				title="YouTube video player"
				frameborder="0"
				allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
				allowfullscreen
			/>
		</q-card-section>

		<q-card-section>
			<div class="text-h6">{{ video.title }}</div>
			<q-markdown
				no-heading-anchor-links
				:src="`${video.description}`"
			></q-markdown>
		</q-card-section>

		<q-card-section
			v-if="video.meta !== null && video.meta.description !== null"
		>
			<q-separator />
			<div class="text-h6">
				{{ $t('components.video-details.description.title') }}
			</div>
			<q-markdown
				no-heading-anchor-links
				:src="`${video.meta.description[i18n.locale.value as 'de-DE' | 'en-US']}`"
			></q-markdown>
		</q-card-section>

		<q-card-section
			v-if="video.meta !== null && video.meta.correction !== null"
		>
			<q-separator />
			<q-banner dense class="warning-banner">
				<div class="text-h6">
					{{ $t('components.video-details.correction.title') }}
				</div>
				<q-markdown
					no-heading-anchor-links
					:src="`${video.meta.correction[i18n.locale.value as 'de-DE' | 'en-US']}`"
				></q-markdown>
			</q-banner>
		</q-card-section>

		<q-card-actions>
			<q-btn
				flat
				color="primary"
				label="text sources"
				:icon="
					showSources ? 'keyboard_arrow_up' : 'keyboard_arrow_down'
				"
				@click="showSources = !showSources"
			/>
			<q-space />
			<q-btn
				flat
				color="primary"
				label="open"
				:href="video.url"
				target="_blank"
			/>
		</q-card-actions>

		<q-slide-transition>
			<q-card-section v-show="showSources" class="q-pt-none">
				<q-separator />
				<q-tabs
					v-model="selectedSource"
					dense
					class="text-grey"
					active-color="primary"
					indicator-color="primary"
					align="justify"
					narrow-indicator
				>
					<q-tab name="text" label="Text" />
					<q-tab name="video" label="Video" />
					<q-tab name="audio" label="Audio" />
				</q-tabs>

				<q-separator />

				<q-tab-panels v-model="selectedSource" animated>
					<q-tab-panel name="text">
						<div class="text-h6">Mails</div>
						Lorem ipsum dolor sit amet consectetur adipisicing elit.
					</q-tab-panel>

					<q-tab-panel name="video">
						<div class="text-h6">Alarms</div>
						Lorem ipsum dolor sit amet consectetur adipisicing elit.
					</q-tab-panel>

					<q-tab-panel name="audio">
						<div class="text-h6">Movies</div>
						Lorem ipsum dolor sit amet consectetur adipisicing elit.
					</q-tab-panel>
				</q-tab-panels>
			</q-card-section>
		</q-slide-transition>
	</q-card>
</template>

<script lang="ts">
import { VideoType } from 'src/types/videos';
import { defineComponent, PropType, ref } from 'vue';
import { useI18n } from 'vue-i18n';

export default defineComponent({
	name: 'VideosPage',
	props: {
		video: {
			type: Object as PropType<VideoType>,
			required: true,
		},
	},
	setup() {
		const stars = ref(0);
		const showSources = ref(false);
		const selectedSource = ref('text');
		const i18n = useI18n();

		return {
			stars,
			showSources,
			selectedSource,
			i18n,
		};
	},
});
</script>
