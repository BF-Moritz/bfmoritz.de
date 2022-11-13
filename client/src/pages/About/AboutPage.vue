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
						{{ $t('pages.about.title') }}
					</div>
				</div>
			</template>
		</q-parallax>
		<div class="container">
			<q-timeline :layout="layout" color="secondary">
				<q-timeline-entry title="Geburt" subtitle="23.Juli 1999">
					<div>Geburt</div>
				</q-timeline-entry>

				<q-timeline-entry title="Grundschule" subtitle="2006 - 2010">
					<div>Grundschule</div>
				</q-timeline-entry>
				<q-timeline-entry title="Gymnasium" subtitle="2010 - 2018">
					<div>Gymnasium</div>
				</q-timeline-entry>

				<q-timeline-entry title="Studium" subtitle="2018 - 2020">
					<div>Informatik Studium am KIT in Karlsruhe</div>
				</q-timeline-entry>

				<q-timeline-entry title="Ausbildung" subtitle="2020 - Jetzt">
					<div>
						Ausbildung zum Fachinforatiker für Anwendungsentwicklung
						bei Tecxipio
					</div>
				</q-timeline-entry>
			</q-timeline>
		</div>
	</q-page>
</template>

<script lang="ts">
import { useQuasar } from 'quasar';
import { computed, defineComponent } from 'vue';

export default defineComponent({
	name: 'AboutPage',
	components: {},
	setup() {
		const $q = useQuasar();

		return {
			layout: computed(() => {
				return $q.screen.lt.md ? 'dense' : 'comfortable';
			}),
		};
	},
});
</script>
