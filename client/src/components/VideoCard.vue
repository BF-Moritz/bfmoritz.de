<template>
	<q-card square flat bordered>
		<q-card-section horizontal class="q-pa-sm">
			<q-img class="col-5" :src="video.thumbnail.url" />
			<div class="col-7 column q-px-md">
				<h5 class="text-h5 q-my-sm">
					{{ video.title }}
				</h5>
				<h6 class="text-h6 q-my-xs">
					{{ video.channel_title }}
				</h6>
				<div class="q-mb-auto">
					<q-chip
						dense
						clickable
						:class="{
							'tag-selected': selectedTagsMap[tag] === true,
							'tag-not-selected': selectedTagsMap[tag] !== true,
						}"
						v-for="tag in video.tags"
						:key="tag"
						@click="setTag(tag)"
					>
						{{ tag }}
					</q-chip>
				</div>
				<q-card-actions align="right" class="self-end">
					<q-btn flat @click="openFull">
						{{ $t('pages.videos.open') }}
					</q-btn>
				</q-card-actions>
			</div>
		</q-card-section>
	</q-card>
</template>
<script lang="ts">
import { computed, defineComponent, PropType } from 'vue';
import { VideoType } from 'src/types/videos';
import { TagMapType } from 'src/types/tag';

export default defineComponent({
	name: 'VideoCard',
	props: {
		video: {
			type: Object as PropType<VideoType>,
			required: true,
		},
		selectedTags: {
			type: Array<string>,
			required: true,
		},
	},
	setup(props, { emit }) {
		function setTag(tag: string) {
			emit('set-tag', tag);
		}

		function openFull() {
			emit('open-full');
		}

		const selectedTagsMap = computed(() => {
			return props.selectedTags.reduce((acc, tag) => {
				acc[tag] = true;
				return acc;
			}, {} as TagMapType);
		});

		return {
			setTag,
			openFull,
			selectedTagsMap,
		};
	},
});
</script>
