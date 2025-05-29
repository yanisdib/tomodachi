import type { Config } from 'orval';

const config: Config = {
    tomodachiAPI: {
        input: './openapi/tomodachi.yaml',
        output: {
            target: './src/api/tomodachi/',
            schemas: './src/api/tomodachi/schemas',
            client: "react-query",
            mode: 'tags-split',
        },
    },
};

export default config;