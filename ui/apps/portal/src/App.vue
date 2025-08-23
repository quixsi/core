<script setup lang="ts">
import { RouterView } from 'vue-router'
import { Toast } from '@quixsi/components'
import { useToast } from '@/composables/useToast'

const { toasts, removeToast } = useToast()
</script>

<template>
  <Toast :toasts="toasts" @remove="removeToast" />
  <RouterView />
</template>
