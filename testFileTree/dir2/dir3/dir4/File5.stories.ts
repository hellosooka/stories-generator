import File5 from './File5.vue'
import { Meta, Story } from "@storybook/vue3";
import { defineComponent } from "vue";

export default {
  title: "story/File5",
  component: File5,
  argTypes: {
  },
  parameters: {
    docs: {
      source: {
        code: [
          "import { File5 } from '@estelink/story'\n\n",
          "<File5 />",
        ].join("\n"),
      },
    },
  },
};

const Template: Story = (args) =>
  defineComponent({
    components: { File5 },
    setup: () => ({ args }),
    template: `<File5 v-bind="args" />`,
  });

export const Default = Template.bind({});
