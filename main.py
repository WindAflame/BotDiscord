import feedparser
from discord.ext import commands

# Créez un nouveau bot Discord en utilisant votre token d'authentification
bot = commands.Bot(command_prefix='!')

# Définissez une fonction pour traiter les nouveaux articles du flux RSS
@bot.event
async def on_ready():
    # Récupérez le canal Discord où les messages seront publiés
    channel = bot.get_channel("CHANNEL_ID")

    # Créez un nouvel objet feedparser en fournissant l'URL du flux RSS
    feed = feedparser.parse("http://exemple.com/rss")

    # Parcourez les articles du flux et publiez leur titre sur Discord
    for entry in feed.entries:
        await channel.send(entry.title)

# Connectez le bot à Discord en utilisant votre token d'authentification
bot.run("TOKEN")
