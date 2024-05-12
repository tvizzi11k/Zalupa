from django.shortcuts import render

def index(request):
    context = {}
    return render(request, 'index.html', context)


def home(request):
    context = {}
    return render(request, 'home.html', context)
