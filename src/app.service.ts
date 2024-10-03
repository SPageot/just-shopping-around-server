import * as dotenv from 'dotenv';
import { Injectable } from '@nestjs/common';
import * as NewsAPI from 'newsapi';
dotenv.config();

@Injectable()
export class AppService {
  private readonly newsAPIKey = process.env.NEWS_API_KEY;

  async getNews() {
    const NEWS_API = new NewsAPI(this.newsAPIKey);
    const newsResponse = await NEWS_API.v2.everything({
      q: 'groceries',
      language: 'en',
      pageSize: 30,
      page: 1,
    });
    return newsResponse.articles;
  }
}
