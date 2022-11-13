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
		<div class="container">
			<p style="font-size: 24px">
				Occaecat amet dolore duis enim eu cupidatat aliquip minim
				cupidatat dolore quis proident deserunt consequat. Aute aliquip
				occaecat laborum culpa anim est ex mollit veniam. Lorem amet
				nostrud officia aliqua Lorem culpa nulla exercitation tempor est
				incididunt labore elit. Do cupidatat anim sint consequat eu ex.
				Sit duis sit cillum anim pariatur quis laboris et labore.
				Proident nostrud magna irure deserunt eu laborum. Pariatur velit
				qui consectetur commodo consequat reprehenderit dolor. Cupidatat
				qui est fugiat commodo aute pariatur consectetur consectetur
				duis non adipisicing quis deserunt ut. Laboris pariatur magna
				dolor consequat. Excepteur sit consequat esse proident fugiat
				dolor et ad laborum cupidatat. Est nisi aliquip adipisicing ea
				deserunt est occaecat commodo deserunt consequat velit. Dolore
				enim do amet occaecat ipsum id dolor velit eiusmod labore
				commodo pariatur magna velit. Exercitation adipisicing in aute
				sunt elit. Do consectetur quis voluptate minim cillum magna
				officia enim consectetur culpa magna. Commodo nostrud do sint
				reprehenderit do aliqua. Irure dolore mollit ut anim minim culpa
				velit id reprehenderit. Magna nostrud irure eu est cillum
				consectetur enim tempor. Occaecat aliqua magna elit enim magna
				occaecat adipisicing fugiat pariatur exercitation sint irure
				veniam. Excepteur non veniam consectetur tempor ex in enim
				mollit magna excepteur. Qui ad voluptate do dolor incididunt
				commodo eu ad. Nostrud veniam velit minim qui velit nulla ad id
				tempor nisi ex. Eu nostrud eiusmod veniam ad exercitation enim
				veniam occaecat et amet cupidatat culpa laborum laborum. Cillum
				proident sint do irure sit aute ut elit magna ipsum et mollit
				eu. Pariatur consequat consectetur reprehenderit sunt nostrud
				aliqua quis ad reprehenderit minim aute veniam eu. Anim ex
				incididunt quis aliqua non. Cillum culpa dolor id deserunt sint.
				Pariatur commodo pariatur veniam ipsum pariatur reprehenderit do
				mollit ea elit deserunt. Non laboris sunt exercitation velit
				officia laborum ullamco ad sit anim occaecat ipsum enim. Culpa
				adipisicing elit sit dolor amet quis fugiat. Labore
				reprehenderit ut velit minim. Incididunt veniam tempor esse ea
				laboris sunt nostrud Lorem magna sit aute voluptate laborum. Qui
				commodo laborum id officia dolor minim reprehenderit irure
				occaecat duis fugiat. Ea aute irure Lorem officia esse excepteur
				veniam eiusmod ea irure ex qui aliqua eu. Officia voluptate
				incididunt minim nostrud ex proident commodo aliquip in ea.
				Reprehenderit culpa pariatur reprehenderit velit voluptate nulla
				minim eu qui dolor labore id ex sint. Officia et anim cillum
				deserunt. Dolore excepteur voluptate consectetur velit nisi
				dolor id proident qui. Aliqua elit aliquip ullamco tempor ut ad.
				Cillum elit voluptate excepteur reprehenderit voluptate nostrud
				ut non veniam eiusmod tempor deserunt ipsum. Est non duis velit
				dolor eu dolore adipisicing amet mollit aliquip fugiat. Esse
				ullamco quis dolor labore anim sunt proident exercitation minim
				aute sint. Velit commodo proident esse nisi laboris eiusmod
				tempor. Adipisicing in minim sit velit aliqua ad occaecat
				laboris. Id reprehenderit magna aute sint laborum exercitation
				voluptate ex ut cillum sunt deserunt. Velit culpa cupidatat ad
				est elit excepteur sunt nulla deserunt cupidatat nostrud anim.
				Consequat velit sint ut cupidatat ea fugiat enim esse Lorem
				dolor. Consequat quis in eu ex cupidatat pariatur ipsum. Minim
				et nostrud enim in minim in ut duis reprehenderit. In culpa in
				officia labore minim nisi. Eiusmod commodo est ex fugiat non
				quis labore voluptate pariatur magna dolore commodo dolore.
				Eiusmod laborum mollit velit in amet. Eu aliquip minim esse
				magna occaecat nulla magna eiusmod consequat adipisicing. Est ea
				in incididunt tempor veniam qui consectetur est aliquip. Eu
				aliqua irure occaecat officia reprehenderit laboris. Veniam
				aliquip ipsum ad ea aute. Ut aliquip qui magna ut ex anim
				excepteur Lorem incididunt laborum labore consectetur cupidatat
				sint. Commodo sit mollit do est duis pariatur anim ea elit
				cupidatat commodo culpa occaecat. Magna ex dolore voluptate
				velit elit non proident do do nisi sunt. Excepteur adipisicing
				magna cupidatat sit nisi velit aliqua in non labore deserunt
				consectetur tempor adipisicing. Ea amet incididunt est amet quis
				sint veniam adipisicing in in sint elit ullamco. Aliqua nostrud
				sit do amet ad eiusmod fugiat do qui nulla incididunt aute esse.
				Non fugiat occaecat consectetur dolor irure eiusmod laborum
				excepteur. Aute culpa officia consequat incididunt sint
				consequat non mollit eu irure reprehenderit sit tempor eu.
				Aliqua elit deserunt veniam velit ea veniam ipsum amet. Esse do
				incididunt excepteur eu veniam deserunt proident cillum
				excepteur incididunt quis. Nostrud Lorem velit ea labore
				proident amet esse quis aliqua id commodo. Cupidatat commodo in
				cillum voluptate esse esse proident laboris reprehenderit
				laboris ut. Duis qui velit eu cupidatat veniam eu incididunt
				deserunt laborum amet. Laborum veniam enim sint cillum
				incididunt officia non. Qui nulla adipisicing mollit duis velit
				cillum dolore laboris anim officia fugiat est. Duis veniam Lorem
				non commodo nostrud duis ut sint. Magna est in commodo pariatur
				est proident excepteur commodo pariatur sint duis. Elit nostrud
				officia reprehenderit esse. Aliqua velit est excepteur irure.
			</p>
		</div>
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
