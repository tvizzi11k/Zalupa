from telegram import Update, InlineKeyboardButton, InlineKeyboardMarkup, WebAppInfo
from telegram.ext import Application, CommandHandler, ContextTypes

TOKEN = '7120451302:AAFmiJriqQlj0aDohxsBjIRdTQbbJOMEyhs'

async def start(update: Update, context: ContextTypes.DEFAULT_TYPE):
    keyboard =[
        [InlineKeyboardButton("голые сиськи", web_app=WebAppInfo(url="http://my.porno365.expert/"))]
    ]
    reply_markup = InlineKeyboardMarkup(keyboard)
    await update.message.reply_text("порно видео секс", reply_markup=reply_markup)

def setup_bot():
    application = Application.builder().token(TOKEN).build()
    application.add_handler(CommandHandler("start", start))
    return application

app = setup_bot()

if __name__ == '__main__':
    app.run_polling()