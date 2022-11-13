import { boot } from 'quasar/wrappers';
import axios, { AxiosInstance } from 'axios';

declare module '@vue/runtime-core' {
	interface ComponentCustomProperties {
		$axios: AxiosInstance;
	}
}

const api = axios.create({ baseURL: 'http://localhost:6969/api/v1' });

export default boot(({ app }) => {
	// for use inside Vue files (Options API) through this.$axios and this.$api
	app.config.globalProperties.$axios = axios;
	app.config.globalProperties.$api = api;
});

export { api };
