<template>
	<q-item-section v-if="icon" avatar>
		<q-icon :name="icon" />
	</q-item-section>

	<q-item-section>
		<q-item-label>{{ title }}</q-item-label>
		<q-item-label caption>{{ caption }}</q-item-label>
	</q-item-section>

	<q-item-section side v-if="external">
		<q-icon name="mdi-open-in-new" />
	</q-item-section>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
	name: 'FlatButton',
	props: {
		title: {
			type: String,
			required: true,
		},

		caption: {
			type: String,
			default: '',
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
});
</script>
