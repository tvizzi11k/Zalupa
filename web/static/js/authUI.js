const tonConnectUI = new TON_CONNECT_UI.TonConnectUI({
  manifestUrl: 'https://176-99-11-185.cloudvps.regruhosting.ru/tonconnect-manifest.json',
});

const unsubscribe = tonConnectUI.onStatusChange(state => {
  window.location.replace('/home')
});

// const unsubscribeModal = tonConnectUI.onModalStateChange((state) => {
//   const currentModalState = tonConnectUI.modalState;
//   console.log(currentModalState)
//   }
// );

// function ChangeBtnStyles(){
//   let btn = document.querySelector('[data-tc-button="true"]');
// }

async function modal() {
  await tonConnectUI.openModal();
}

const unsubscribeWallet = tonConnectUI.onStatusChange(
  walletAndwalletInfo => {
      console.log(walletAndwalletInfo)
  } 
);
