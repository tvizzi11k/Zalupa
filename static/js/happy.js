const tonConnectUI = new TON_CONNECT_UI.TonConnectUI({
    manifestUrl: 'https://176-99-11-185.cloudvps.regruhosting.ru/static/ton.json',
    buttonRootId: 'ton-connect',
});

const unsubscribe = tonConnectUI.onStatusChange(state => {
    window.location.replace('/home')
});

tonConnectUI.uiOptions = {
  uiPreferences: {
    borderRadius: 'm',
    theme: 'DARK',
    }
}