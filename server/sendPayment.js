var request = require('request');
var StellarSdk = require('stellar-sdk');

request.post(
  {
    url: 'http://localhost:8006/payment',
    form: {
      amount: '1',
      asset: StellarSdk.Asset.native(),
      destination: 'GB5FUITDOLTZDQ2YDZBRRDIHE56RLD5TXRSI4FESFN2E74EBLLOHE6YN',
      source: 'SCSJFTW75JRUKTGPABTGNHKKLR2QTANJPID4RHDVZWRFTSFPGHY2GIXJ'
    }
  },
  function(error, response, body) {
    if (error || response.statusCode !== 200) {
      console.error('ERROR!', error || body);
    } else {
      console.log('SUCCESS!', body);
    }
  }
);
