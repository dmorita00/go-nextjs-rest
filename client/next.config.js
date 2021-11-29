module.exports = {
  reactStrictMode: true,
  env: {
    ...require(`./config/${process.env.NODE_ENV || 'local'}.json`),
  }
}
