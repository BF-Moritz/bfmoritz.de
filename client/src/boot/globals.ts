import { boot } from 'quasar/wrappers';

export const vars = {
	url: 'bfmoritz.de',
};

export default boot(({ app }) => {
	app.config.globalProperties.$vars = vars;
});
