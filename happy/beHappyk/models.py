from django.db import models

# Create your models here.
class User(models.Model):
    username = models.CharField(max_length=255)
    balance = models.DecimalField(max_digits=10, decimal_places=2, default=0)

class Promocode(models.Model):
    code = models.CharField(max_length=255, unique=True)
    value = models.DecimalField(max_digits=10, decimal_places=2)
    used = models.BooleanField(default=False)
