// буду дописывать
const tonConnectUI = new TON_CONNECT_UI.TonConnectUI({
  // manifestUrl: 'https://176-99-11-185.cloudvps.regruhosting.ru/tonconnect-manifest.json',
  manifestUrl: 'https://github.com/tvizzi11k/Zalupa/tonconnect-manifest.json',
});

tonConnectUI.onStatusChange(console.log)

const unsubscribe = tonConnectUI.onStatusChange(state => {
  window.location.replace('/home')
});

const unsubscribeModal = tonConnectUI.onModalStateChange((state) => {
  const currentModalState = tonConnectUI.modalState;
  console.log(currentModalState)
    if(currentModalState['status'] == "opened"){
      console.log("good")
      ChangeBtnStyles()
    } else{
      console.log("bad")
    }
  }
);

function ChangeBtnStyles(){
  let btn = document.querySelector('[data-tc-button="true"]');
}

async function modal() {
  await tonConnectUI.openModal();
}
