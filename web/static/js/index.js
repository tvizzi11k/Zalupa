Autorization()
function Autorization(){
  if (localStorage.getItem('ton-connect-storage_bridge-connection') === null) {
    window.location.replace('/')
  }
}

const tonConnectUI = new TON_CONNECT_UI.TonConnectUI({
  manifestUrl: 'https://176-99-11-185.cloudvps.regruhosting.ru/tonconnect-manifest.json',
});

async function disconect(){
  await tonConnectUI.disconnect();
  window.location.replace('/')
}

function getUserKey(){
 const userkey =  JSON.parse(localStorage.getItem('ton-connect-storage_bridge-connection'));
 const key = userkey.session.walletPublicKey;
 return key;
}

// balance text element id
const balanceTextId = "balance-counter";

/**
 * Func for fetch balance
 * @param {{ token: string }} params
 */
async function fetchDashboardData(params) {
  try {
    /* Fetch Balance */

    /**
     * @type {{balance: number}}
     */
    const response = await fetch(
      "https://176-99-11-185.cloudvps.regruhosting.ru/get-balance", // URL ENDPOINT TO GET BALANCE
      {
        method: "GET",
        cache: "no-cache",
        headers: new Headers({
          "Content-Type": "application/json",
          Authorization: params.token,
        }),
      }
    ).then((response) => response.json());

    /* Update Balance DOM */

    const balanceText = document.getElementById(balanceTextId);

    balanceText.textContent = Number(response.balance).toLocaleString("ru-RU");
  } catch (error) {
    console.error(error);

    alert("[DASHBOARD#FETCH_DATA]: Unknown error");
  }
}

/**
 * Func for redeem promo code
 * @param {{ token: string; promoCode: string }} params
 */
async function redeemPromoCode(params) {
  try {
    /* POST Promo Code */

    await fetch(
      "https://176-99-11-185.cloudvps.regruhosting.ru//apply-promocode", // URL ENDPOINT TO GET BALANCE
      {
        method: "POST",
        cache: "no-cache",
        headers: new Headers({
          "Content-Type": "application/json",
          Authorization: params.token,
        }),
        body: JSON.stringify({
          code: params.promoCode,
        }),
      }
    );

    await fetchDashboardData({
      token: params.token,
    });
  } catch (error) {
    console.error(error);

    alert("[DASHBOARD#REDEEM_PROMO_CODE]: Unknown error");
  }
}
