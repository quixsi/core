<template>
  <Card>
    <CardHeader
      v-if="$slots.title || $slots.description"
    >
      <CardTitle v-if="$slots.title">
        <slot name="title" />
      </CardTitle>
      <CardDescription
        v-if="$slots.description"
      >
        <slot name="description" />
      </CardDescription>
    </CardHeader>
    <CardContent
      v-if="$slots.content"
    >
      <slot name="content" />
    </CardContent>
    <CardFooter
      v-if="$slots.footer"
    >
      <slot name="footer" />
    </CardFooter>
  </Card>
</template>

<script setup lang="ts">
import Card from './Card.vue';
import CardContent from './CardContent.vue';
import CardDescription from './CardDescription.vue';
import CardFooter from './CardFooter.vue';
import CardHeader from './CardHeader.vue';
import CardTitle from './CardTitle.vue';
</script>
