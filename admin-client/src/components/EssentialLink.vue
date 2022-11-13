<template>
	<q-item clickable tag="a" target="_blank" :href="link">
		<flat-button
			:title="title"
			:caption="caption"
			:icon="icon"
			:external="external"
		/>
	</q-item>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import FlatButton from './FlatButton.vue';

export default defineComponent({
	name: 'EssentialLink',
	props: {
		title: {
			type: String,
			required: true,
		},
		caption: {
			type: String,
			default: '',
		},
		link: {
			type: String,
			default: '#',
		},
		icon: {
			type: String,
			default: '',
		},
		external: {
			type: Boolean,
			default: false,
		},
	},
	components: { FlatButton },
});
</script>
