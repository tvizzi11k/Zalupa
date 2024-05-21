Autorization()

function Autorization(){
  if (localStorage.getItem('ton-connect-storage_bridge-connection') === null) {
    window.location.replace('/')
  }
}

const tonConnectUI = new TON_CONNECT_UI.TonConnectUI({
  manifestUrl: 'https://176-99-11-185.cloudvps.regruhosting.ru/static/tonconnect-manifest.json',
  // manifestUrl: 'https://github.com/tvizzi11k/Zalupa/tonconnect-manifest.json',
});

tonConnectUI.onStatusChange(console.log)

const unsubscribe = tonConnectUI.onStatusChange(state => {
  window.location.replace('/home')
});

const unsubscribeModal = tonConnectUI.onModalStateChange((state) => {
  const currentModalState = tonConnectUI.modalState;
  console.log(currentModalState)
  }
);

function ChangeBtnStyles(){
  let btn = document.querySelector('[data-tc-button="true"]');
}

async function modal() {
  await tonConnectUI.openModal();
}
