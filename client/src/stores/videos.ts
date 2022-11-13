import { api } from 'src/boot/axios';
import { VideoType } from 'src/types/videos';

export class VideosStore {
	public static async getVideos(): Promise<null | VideoOutType> {
		try {
			const { data, status } = await api.get('/videos');
			if (status != 200) return null;
			return { timestamp: new Date(data.timestamp), videos: data.videos };
		} catch (err) {
			console.error(err);
			return null;
		}
	}
}

type VideoOutType = {
	timestamp: Date;
	videos: VideoType[];
};
