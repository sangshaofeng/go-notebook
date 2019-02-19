module.exports = {
  baseUrl: './',
  productionSourceMap: false,
  devServer: {
    proxy: {
      '*': {
        target: 'http://127.0.0.1:8080',
        changeOrigin: true,
        cookieDomainRewrite: "localhost"
      },
    }
  },
}