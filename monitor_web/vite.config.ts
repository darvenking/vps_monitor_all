import {splitVendorChunkPlugin, UserConfigExport} from 'vite'
import vue from '@vitejs/plugin-vue2'
import legacy from '@vitejs/plugin-legacy'
import Components from 'unplugin-vue-components/vite'
import {ElementUiResolver} from 'unplugin-vue-components/resolvers'
import path from 'path'

export default (): UserConfigExport => {

  return {
    server: {
      host: true,
      port: 8080,
      proxy: {
        '/api': {
          target: 'http://127.0.0.1:8000',
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api/, '')
        }
      },
    },
    plugins: [
      vue(),
      splitVendorChunkPlugin(),
      legacy({
        targets: ['defaults', 'not ie < 9'],
      }),
      Components({
        resolvers: [ElementUiResolver({
          importStyle: true,
          }
        )],
      }),
    ],
    build: {
      target: 'es2015',
      chunkSizeWarningLimit: 2000,
      rollupOptions: {
        output: {
          manualChunks: {
            'element-ui': ['element-ui'],
          },
        },
      },
    },
    resolve: {
      // 配置路径别名
      // alias: {
      //   '@': path.resolve(__dirname, './src'),
      // },
      alias: [
        {
          find: /\/@\//,
          replacement: path.resolve('src') + '/',
        },
        {
          find: /@\//,
          replacement: path.resolve('src') + '/',
        },
      ],
    },
  }
}
