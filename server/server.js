const ngrok = require('ngrok');

(async function ngrokInit() {
  const url = await ngrok.connect({
    proto: 'tcp',
    addr: 8080,
    authtoken: '21HFfNCHfamQZRYAmDfSUKooqvH_4x2u6MwUMx6neEcpsn8Bz'
  })
  console.log("game url: ", url)
})()
