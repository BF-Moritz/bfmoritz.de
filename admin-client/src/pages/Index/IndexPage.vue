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
					<div class="text-h1 text-white text-center">BF_Moritz</div>
				</div>
			</template>
		</q-parallax>
		<div class="container"></div>
	</q-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
	name: 'IndexPage',
	components: {},
	setup() {
		return {};
	},
});
</script>
