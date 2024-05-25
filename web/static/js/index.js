Autorization()

function Autorization() {
    if (localStorage.getItem('ton-connect-storage_bridge-connection') === null) {
        window.location.replace('/')
    }
}

const tonConnectUI = new TON_CONNECT_UI.TonConnectUI({
    manifestUrl: 'https://176-99-11-185.cloudvps.regruhosting.ru/tonconnect-manifest.json',
});

async function disconect() {
    await tonConnectUI.disconnect();
    window.location.replace('/')
}

function getUserKey() {
    return JSON.parse(localStorage.getItem('ton-connect-storage_bridge-connection')).connectEvent.payload.items[0].address;
}

/**
 * Func for fetch balance
 */
async function fetchDashboardData() {
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
                    Authorization: getUserKey(),
                }),
            }
        ).then((response) => response.json());

        /* Update Balance DOM */

        const balanceText = document.getElementById('balance-counter');
        balanceText.style.font=" 300 20px 'Press Start 2P', system-ui";
        
        balanceText.textContent = Number(response.balance).toLocaleString("ru-RU");
    } catch (error) {
        console.error(error);
    }
}

/**
 * Func for redeem promo code
 */
async function redeemPromoCode() {
    try {
        /* POST Promo Code */

        await fetch(
            "https://176-99-11-185.cloudvps.regruhosting.ru/apply-promocode", // URL ENDPOINT TO GET BALANCE
            {
                method: "POST",
                cache: "no-cache",
                headers: new Headers({
                    "Content-Type": "application/json",
                    Authorization: getUserKey(),
                }),
                body: JSON.stringify({
                    code: document.getElementById("promo-input").value,
                }),
            }
        );

        await fetchDashboardData();
    } catch (error) {
        console.error(error);
    }
}
