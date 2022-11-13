<template>
	<q-page ref="page" class="column justify-start items-center">
		<q-parallax :height="600" :speed="1">
			<template v-slot:media>
				<img
					src="https://www.zdf.de/assets/bg-landscape-photography-award-110~2400x1350?cb=1637152668127"
				/>
			</template>

			<template v-slot:content="scope">
				<div
					class="absolute column items-center"
					:style="{
						opacity: (1 - scope.percentScrolled) * 4 - 1,
						top: scope.percentScrolled * 80 + '%',
						left: 0,
						right: 0,
					}"
				>
					<div class="text-h1 text-white text-center">
						{{ $t('pages.videos.title') }}
					</div>
				</div>
			</template>
		</q-parallax>
		<div class="container full-width q-px-sm">
			<q-card square flat bordered class="q-mt-md q-pa-sm">
				<!-- Search bar -->
				<q-input
					filled
					:debounce="100"
					v-model="videoSearch"
					:label="$t('pages.videos.search')"
					stack-label
					dense
					square
					:disable="videos === null"
				/>
				<!-- Tags -->
				<div class="tags-container">
					<q-chip
						v-if="videos !== null"
						class="q-mt-sm"
						dense
						clickable
						:class="{
							'q-mt-sm': true,
						}"
						@click="resetTags"
					>
						{{ $t('pages.videos.tags-caption') }}:
						<q-tooltip>{{
							$t('pages.videos.tags-reset')
						}}</q-tooltip>
					</q-chip>
					<div v-else class="q-mt-sm q-ml-sm">
						{{ $t('pages.videos.tags-loading') }}
					</div>
					<q-chip
						class="q-mt-sm"
						dense
						clickable
						:class="{
							'q-mt-sm': true,
							'tag-selected': tag.selected,
							'tag-not-selected': !tag.selected,
						}"
						v-for="tag in tags"
						:key="tag.tag"
						@click="toggleTag(tag.tag)"
					>
						{{ tag.tag }}
					</q-chip>
				</div>
			</q-card>

			<!-- Video List -->
			<div v-if="videos !== null" class="video-list">
				<VideoCard
					v-for="video in filteredVideos"
					:key="video.id"
					:video="video"
					:selected-tags="selectedTags"
					class="q-my-md"
					@set-tag="toggleTag"
					@open-full="openVideo(video)"
				/>
			</div>
			<!-- TODO -->
			<div v-else>Loading...</div>
			<q-dialog v-model="showVideoDetails">
				<VideoDetails
					v-if="selectedVideo !== null"
					:video="selectedVideo"
				/>
			</q-dialog>
		</div>
	</q-page>
</template>

<script lang="ts">
import { defineComponent, onMounted, Ref, ref, computed } from 'vue';
import { VideoType } from 'src/types/videos';
import { TagType } from 'src/types/tag';
import { VideosStore } from 'src/stores/videos';

import VideoCard from 'src/components/VideoCard.vue';
import VideoDetails from '../../components/VideoDetails.vue';

export default defineComponent({
	name: 'VideosPage',
	components: { VideoCard, VideoDetails },
	setup() {
		const videos: Ref<null | VideoType[]> = ref(null);
		const lastUpdated: Ref<null | Date> = ref(null);
		const tags: Ref<TagType[]> = ref([]);
		const videoSearch: Ref<string> = ref('');
		const selectedVideo: Ref<null | VideoType> = ref(null);
		const showVideoDetails: Ref<boolean> = ref(false);

		const selectedTags = computed(() => {
			return tags.value
				.filter((tag) => tag.selected)
				.map((tag) => tag.tag);
		});

		const filteredVideos = computed(() => {
			if (videos.value === null) return null;
			const v = videos.value.filter((vid) => {
				if (selectedTags.value.length > 0) {
					let hasTag = false;
					for (const tag of vid.tags) {
						if (selectedTags.value.includes(tag)) {
							hasTag = true;
						}
					}
					if (!hasTag) return false;
				}

				if (videoSearch.value.length > 0) {
					if (
						!vid.title
							.toLowerCase()
							.includes(videoSearch.value.trim().toLowerCase())
					) {
						return false;
					}
				}

				return true;
			});
			return v;
		});

		function toggleTag(tag: string) {
			const allTags = tags.value;
			const i = allTags.findIndex((t) => t.tag === tag);
			if (i >= 0) {
				allTags[i].selected = !allTags[i].selected;
			}
			tags.value = allTags;
		}

		function resetTags() {
			const allTags = tags.value;
			for (const tag of allTags) {
				tag.selected = false;
			}
			tags.value = allTags;
		}

		function openVideo(video: VideoType) {
			selectedVideo.value = video;
			showVideoDetails.value = true;
		}

		function closeVideo() {
			showVideoDetails.value = false;
			selectedVideo.value = null;
		}

		onMounted(async () => {
			const result = await VideosStore.getVideos();
			if (result === null) {
				return;
			}

			videos.value = result.videos;
			lastUpdated.value = result.timestamp;

			const tempTags = new Set<string>();

			for (const video of videos.value) {
				for (const tag of video.tags) {
					tempTags.add(tag);
				}
			}

			tags.value = Array.from(tempTags).map((t) => {
				return { tag: t, selected: false };
			});
		});

		return {
			videos,
			lastUpdated,
			tags,
			videoSearch,
			selectedTags,
			filteredVideos,
			selectedVideo,
			toggleTag,
			resetTags,
			openVideo,
			closeVideo,
			showVideoDetails,
		};
	},
});
</script>

<style lang="scss" src="./VideosPage.scss" scoped />
