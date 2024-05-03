from django.shortcuts import render, redirect
from .models import User, Promocode
from django.contrib import messages

# Create your views here.
def index(request):
    return render(request, 'index.html')

def home(request):
    if request.method == 'POST':
        promocode_input = request.POST.get('promocode')
        try:
            promocode = Promocode.objects.get(code=promocode_input)
            if promocode.used:
                messages.error(request, "Промокод использован")
            else:
                user = request.user
                user.balance += promocode.value
                user.save()
                promocode.used = True
                promocode.save()
                messages.success(request, f"Промокод на сумму {promocode.value} применен")
        except Promocode.DoesNotExist:
            messages.error(request, "Промокода не существует.")
    
    return render(request, 'index.html', {'message': messages.get_messages(request)})