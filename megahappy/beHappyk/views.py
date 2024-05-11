from django.shortcuts import render

async def home(request):
    context = {}
    return render(request, 'home.html', context)
