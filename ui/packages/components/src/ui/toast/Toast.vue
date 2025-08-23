<script setup lang="ts">
import { ref, watch, onUnmounted } from 'vue'
import type { HTMLAttributes } from 'vue'
import { cn } from '#lib/utils'

const props = defineProps<{
    message: string
    type?: 'success' | 'error'
    duration?: number
    class?: HTMLAttributes['class']
    toastKey?: string | number
}>()

const visible = ref(false)
let timeout: ReturnType<typeof setTimeout> | null = null

watch(() => props.toastKey, (newKey) => {
    if (timeout) {
        clearTimeout(timeout)
    }

    visible.value = false

    if (!props.message) {
        return
    }

    setTimeout(() => {
        visible.value = true
        timeout = setTimeout(() => {
            visible.value = false
        }, props.duration ?? 3000)
    }, 10)
}, { immediate: true })

onUnmounted(() => {
    if (timeout) {
        clearTimeout(timeout)
    }
})
</script>

<template>
    <transition name="fade">
        <div
            v-if="visible"
            :key="props.toastKey"
            data-slot="toast"
            :class="cn(
                'fixed z-50 top-8 right-8 px-4 py-2 rounded shadow-lg',
                type === 'error' ? 'bg-red-500 text-white' : 'bg-green-500 text-white',
                props.class,
            )"
        >
            {{ message }}
        </div>
    </transition>
</template>

<style scoped>
.fade-enter-active, .fade-leave-active {
    transition: opacity 0.3s;
}

.fade-enter-from, .fade-leave-to {
    opacity: 0;
}
</style>
