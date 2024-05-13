const connector = new TonConnectSDK.TonConnect({
    manifestUrl: 'https://176-99-11-185.cloudvps.regruhosting.ru/static/ton.json'
});

connector.restoreConnection();

connector.onStatusChange(walletInfo => {
    console.log(walletInfo)
    window.location.replace('/home')
})

function auth() {
    const walletConnectionSource = {
        universalLink: 'https://app.tonhub.com/ton-connect',
        bridgeUrl: 'https://bridge.tonapi.io/bridge'
    }
    
    const universalLink = connector.connect(walletConnectionSource);
    
    window.open(universalLink, '_blank')
}
// universalLink: 'https://t.me/wallet',