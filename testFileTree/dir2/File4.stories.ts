import File4 from './File4.vue'
import { Meta, Story } from "@storybook/vue3";
import { defineComponent } from "vue";

export default {
  title: "test/File4",
  component: File4,
  argTypes: {
  },
  parameters: {
    docs: {
      source: {
        code: [
          "import { File4 } from '@estelink/test'\n\n",
          "<File4 />",
        ].join("\n"),
      },
    },
  },
};

const Template: Story = (args) =>
  defineComponent({
    components: { File4 },
    setup: () => ({ args }),
    template: `<File4 v-bind="args" />`,
  });

export const Default = Template.bind({});
