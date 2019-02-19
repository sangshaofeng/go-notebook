module.exports = {
    devServer: {
        proxy: {
          '/': {
            target: 'localhost:8000',
            changeOrigin: true
          },
        }
      }
  }