from telegram import Update, InlineKeyboardButton, InlineKeyboardMarkup, WebAppInfo
from telegram.ext import Application, CommandHandler, ContextTypes

TOKEN = '7120451302:AAFmiJriqQlj0aDohxsBjIRdTQbbJOMEyhs'

async def start(update: Update, context: ContextTypes.DEFAULT_TYPE):
    keyboard =[
        [InlineKeyboardButton("открыть", web_app=WebAppInfo(url="https://behappyhappyhappy1488.netlify.app/"))]
    ]
    reply_markup = InlineKeyboardMarkup(keyboard)
    await update.message.reply_text("сайт", reply_markup=reply_markup)

def setup_bot():
    application = Application.builder().token(TOKEN).build()
    application.add_handler(CommandHandler("start", start))
    return application

app = setup_bot()

if __name__ == '__main__':
    app.run_polling()